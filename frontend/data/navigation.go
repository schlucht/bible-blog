package data

type Navigation struct {
	Link    string
	Icon    string
	Alt     string
	Content string
	Class   string
}

func NavigationDatas() []Navigation {

	defaultClass := "bg-green-300 text-2xl"

	return []Navigation{
		{
			Link:    "/",
			Icon:    "🛖",
			Alt:     "Home",
			Content: "Startseite",
			Class:   defaultClass,
		},
		{
			Link:    "/about",
			Icon:    "🕵️‍♀️",
			Alt:     "About",
			Content: "Über uns...",
			Class:   defaultClass,
		},
		{
			Link:    "/services",
			Icon:    "😊",
			Alt:     "Services",
			Content: "Wir bieten",
			Class:   defaultClass,
		},
	}
}
