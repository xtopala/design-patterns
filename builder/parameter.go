// The Builder Parameter

// A question could someday cross our minds.
// And we might be asking how do we get the users
// of our API to actually use our builders as opposed to stop
// messing with the object direclty.

// One approach to this is that we simply hide our precious objects,
// that we want to hide from our users clammy fingers.

// Let's suppose that we have an API for sending some kind of emails.

package main

import "strings"

type email struct {
	from, to, subject, body string
}

// Unfortunately, the problem, or at least one of the problems, is that
// we want our emails to be fully specified.

// For this we could write a validator so we can write a component which
// does it job on each email.
// Or what we can do instead is we can create a builder so that the user can
// invoke methods on that builder in order to build that pesky email.
// So we can keep the email struct private and we don't let it kind of bleed
// out from outside the package.

// But what we can do is we can create a more available type that builds emails.
// And it's not going to expose the different parts of the email directly.

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	// and we can do a simple validation
	if !strings.Contains(from, "@") {
		panic("dis not an email foo")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

// So now we come to the important part, and that's how do
// we actually get people to send our email?
// Because somewhere behind the scenes we want to have a function
// which actually takes the email object and it does whatever it need to do.

func sendMailImpl(email *email) {
	//...
}

// But we don't want our clients to actually work with the email object.
// Only with builder.

// Well we can do this by using a -> Builder Parameter.
// That's babically going to be a fucntion which sort of applies to the builder.
// So we have to provide a function which takes a builder ant then does something
// with it, tipically sort of calls something on the builder.

// We're going to use this function in our publicly exposed function called SendEmail.
type build func(*EmailBuilder)

// This function is the function that people are meant to be using.

// -> This is in fact a way we can force declined to use the builder as opposed to
//    providing some sort of incomplete object for initialization.

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("pitty@foo.com").
			To("ateam@baz.com").
			Subject("A-Team").
			Body("Quickly foos")
	})

}
