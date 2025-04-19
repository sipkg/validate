package messages

import "log"

var messages = map[string]map[string]string{
	"fr": {
		"is not a string":                                      "n'est pas une chaîne de catactères",
		"contains non-alphabetic characters":                   "contient des caractères non non-alphabetique",
		"contains non-alphanumeric characters":                 "contient des caractères non-alphanumerique",
		"is not a valid email address":                         "n'est pas une adresse email valide",
		"is not numeric":                                       "n'est pas un nombre",
		"must be greater than %f":                              "doit être supérieur à %f",
		"must be %d characters long":                           "doit avoir %d caractères",
		"must be less than %f":                                 "doit être inférieur à %f",
		"is too long; it must be at most %d characters long":   "est trop long; il doit avoir au plus %d caractères",
		"is too short; it must be at least %d characters long": "est trop court; il doit avoir au moins %d catactères",
		"is empty":                         "est vide",
		"is 0":                             "est zéro",
		"is not a Time type":               "n'est pas de type Time",
		"has a zero value":                 "a une zéro-valeur",
		"doesn't match regular expression": "ne correspond pas à l'expression régulière",
		"is not a valid URL":               "n'est pas une URL valide",
		"has an invalid scheme '%s'":       "a un scheme d'URL invalide '%s'",
		"has an invalid host '%s'":         "a un host d'URL invalide '%s'",
		"is an invalid UUID":               "est un UUID invalide",
	},
}

var lang = "en"

func ChangeLang(l string) {
	lang = l
}

func Translate(msg string) string {
	if lang == "en" {
		return msg
	}
	msg, ok := messages[lang][msg]
	if !ok {
		log.Fatalf("no translation in '%s' for '%s'", lang, msg)
	}
	return msg
}
