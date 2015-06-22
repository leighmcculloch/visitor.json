package main

import (
  "strings"
  "net/http"
  "golang.org/x/text/language"
)

func GetAcceptLanguageResponse(r *http.Request, mr *MainResponse) {
  httpAcceptLanguageHeader := strings.Join(r.Header["Accept-Language"], ",")

  languageTags, _, _ := language.ParseAcceptLanguage(httpAcceptLanguageHeader)

  languages := make([]string, len(languageTags))
  for i := 0; i < len(languages); i++ {
    code := languageTags[i].String()
    languages[i] = code
  }

  mr.AcceptedLanguages = languages
}
