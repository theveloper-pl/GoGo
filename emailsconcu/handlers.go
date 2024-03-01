package main

import (
	"fmt"
	"html/template"
	"main/emailsconcu/data"
	"net/http"
	"strconv"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)


func (app *Config) HomePage(w http.ResponseWriter,r *http.Request) {
    app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter,r *http.Request) {
    app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter,r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)

	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials")
		http.Redirect(w, r, "/login",http.StatusSeeOther)
		return
	}

	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials")
		http.Redirect(w, r, "/login",http.StatusSeeOther)
		return
	}

	if !validPassword{
		msg := Message {
			To: email,
			Subject: "Failed login attempt",
			Data: "Failed login attempt",
		}

		app.sendEmail(msg)

		app.Session.Put(r.Context(), "error", "Invalid credentials")
		http.Redirect(w, r, "/login",http.StatusSeeOther)
		return
	}

	if user.Active == 0 {
		app.Session.Put(r.Context(), "error", "Account is not activated")
		http.Redirect(w, r, "/login",http.StatusSeeOther)
		return
	}
	

	app.Session.Put(r.Context(),"userID", user.ID)
	app.Session.Put(r.Context(),"user", user)

	app.Session.Put(r.Context(), "flash", "Successful login")

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *Config) LogoutPage(w http.ResponseWriter,r *http.Request) {
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter,r *http.Request) {
    app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter,r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	verify_password := r.Form.Get("verify-password")
	first_name := r.Form.Get("first-name")
	last_name := r.Form.Get("last-name")	

	if password != verify_password {
		app.Session.Put(r.Context(), "error", "Password are not matching")
		http.Redirect(w, r, "/register",http.StatusSeeOther)	
		return
	}

	user := data.User{
		Email: email,
		Password: password,
		FirstName: first_name,
		LastName: last_name,
		Active: 0,
		IsAdmin: 0,
	}

	_, err = user.Insert(user)

	if err != nil {
		app.ErrorLog.Println(err)
		app.Session.Put(r.Context(), "error", "Cannot create User !")
		http.Redirect(w, r, "/register",http.StatusSeeOther)	
		return
	}


	url := fmt.Sprintf("http://localhost/activate?email=%s", user.Email)
	signedURL := GenerateTokenFromString(url)

	msg := Message {
		To: user.Email,
		Subject: "Activate your account !",
		Data: template.HTML(signedURL),
		Template: "confirmation-email",
	}

	app.sendEmail(msg)

	app.Session.Put(r.Context(), "flash", "Confirmation email has been sent !")
	http.Redirect(w, r, "/login",http.StatusSeeOther)	
	return

}

func (app *Config) ActivateAccountPage(w http.ResponseWriter,r *http.Request) {
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost%s", url)
	okay :=VerifyToken(testURL)
	if !okay {
		app.Session.Put(r.Context(), "error", "Wrong confirmation URL !")
		http.Redirect(w, r, "/login",http.StatusSeeOther)	
		return
	}

	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found")
		http.Redirect(w, r, "/login",http.StatusSeeOther)	
		return
	}

	u.Active = 1
	err = u.Update()

	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to update user")
		http.Redirect(w, r, "/login",http.StatusSeeOther)	
		return
	}

	app.Session.Put(r.Context(), "flash", "Account activated")
	http.Redirect(w, r, "/login",http.StatusSeeOther)

}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}

func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	
	id := r.URL.Query().Get("id")
	planID, _ := strconv.Atoi(id)

	plan, err := app.Models.Plan.GetOne(planID)
	if err != nil {
		app.Session.Put(r.Context(), "error", "No plan found")
		http.Redirect(w, r, "/members/plans",http.StatusSeeOther)	
		return
	}

	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.Session.Put(r.Context(), "error", "No user found")
		http.Redirect(w, r, "/login",http.StatusSeeOther)	
		return
	}

	app.Wait.Add(1)

	go func() {
		defer app.Wait.Done()
		
		invoice, err := app.getInvoice(user, plan)
		if err != nil {
			app.ErrorChan <- err
		}


		msg := Message{
			To: user.Email,
			Subject: "Invoice",
			Data: invoice,
			Template: "invoice",
		}
		app.sendEmail(msg)


	}()

	app.Wait.Add(1)

	go func ()  {
		defer app.Wait.Done()
		pdf := app.generateManual(user, plan)
		err := pdf.OutputFileAndClose(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))
		if err != nil {
			app.ErrorChan <- err
			return
		}

		msg := Message{
			To: user.Email,
			Subject: "Invoice",
			Data: "Your user manual is attached",
			AttachmentsMap: map[string]string{"Manual.pdf": fmt.Sprintf("./tmp/%d_manual.pdf", user.ID)},
		}
		app.sendEmail(msg)


	}()

	err = app.Models.Plan.SubscribeUserToPlan(user, *plan)
	if err != nil {
		app.Session.Put(r.Context(),"error", "Error subscribing to plan !")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	u, err := app.Models.User.GetOne(user.ID)
	if err != nil {
		app.Session.Put(r.Context(),"error", "Error getting user from Database !")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "user", u)

	app.Session.Put(r.Context(),"flash", "Subscribed !")
	http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
}

func (app *Config) generateManual(u data.User, plan *data.Plan) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)

	importer := gofpdi.NewImporter()


	t := importer.ImportPage(pdf, "./pdf/manual.pdf", 1, "/MediaBox")
	pdf.AddPage()

	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	pdf.SetX(75)
	pdf.SetY(150)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", u.FirstName, u.LastName), "", "C", false)
	pdf.Ln(5)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s User Guide", plan.PlanName), "", "C", false)

	return pdf

}

func (app *Config) getInvoice(u data.User, plan *data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}