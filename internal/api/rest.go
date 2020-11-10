package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	d "github.com/logologics/kunren-be/internal/domain"
	jisho "github.com/logologics/kunren-be/internal/extDict/jisho"
)

// SearchJisho returns the handler for GET /search/jisho/{query}
func (e *Env) SearchJisho(c *gin.Context) {
	query := c.Query("q")

	sr, err := jisho.Search(query)
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Error searching Jisho")
		return
	}

	c.JSON(http.StatusOK, sr)
}

func parseTagsParam(tags string) []string {
	if len(tags) == 0 {
		return []string{}
	}

	return strings.Split(tags, ":")
}

// Vocabs returns all vocabs for the user
func (e *Env) Vocabs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("p"))
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Wrong page param value")
		return
	}
	pageSize, err := strconv.Atoi(c.Query("ps"))
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Wrong page size param value")
		return
	}
	srt, err := d.ParseSorting(c.Query("s"))
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Wrong sorting param value")
		return
	}

	tags := parseTagsParam(c.Query("t"))

	log.Infof("page/pageSize/srt/tags %v/%v/%v/%v", page, pageSize, srt, tags)

	vocabs, err := e.Repo.ListVocabs(page, pageSize, srt, e.User, tags)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err, "Error retrieving Vocabs")
		return
	}

	c.JSON(http.StatusOK, vocabs)
}

// FindVocab returns the vocab with the given key and language
func (e *Env) FindVocab(c *gin.Context) {
	key := c.Query("k")
	lang := c.Query("l")
	check, _ := strconv.ParseBool(c.Query("c"))

	// log.Infof("findVocab key/lang %v/%v", key, lang)

	vocabs, err := e.Repo.FindVocab(e.User, d.ToLanguage(lang), key)
	if err != nil && check {
		sendError(c, http.StatusOK, err, "Not found")
		return
	}

	if err != nil {
		sendError(c, http.StatusNotFound, err, "Not found")
		return
	}

	c.JSON(http.StatusOK, vocabs)
}

// Remember stores a search result in the dict history
func (e *Env) Remember(c *gin.Context) {
	var word d.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		sendError(c, http.StatusBadRequest, err, "Can't bind word")
		return
	}

	storedWord, err := e.Repo.StoreWord(word)
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Can't store word")
		return
	}

	tags := parseTagsParam(c.Query("t"))

	vocab := d.Vocab{
		Key:           storedWord.Key,
		WordID:        storedWord.ID,
		UserID:        e.User.ID,
		Language:      word.Language,
		SearchStrings: []string{storedWord.Lexeme, storedWord.Lemma.Reading},
		Tags:          tags,
	}
	vocab, err = e.Repo.StoreVocab(vocab, true)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err, "Can't create vocab")
		return
	}

	c.JSON(http.StatusOK, vocab)
}

// Tags returns the list of tags of a user
func (e *Env) Tags(c *gin.Context) {

	tags, err := e.Repo.Tags(e.User.ID)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err, "Can't list tags")
		return
	}

	c.JSON(http.StatusOK, d.Tags{Tags: tags})
}

// DeleteTag deletes a tag from all documents of the user of a user
func (e *Env) DeleteTag(c *gin.Context) {
	tag := c.Param("tag")

	if err := e.Repo.DeleteTag(e.User.ID, tag); err != nil {
		sendError(c, http.StatusInternalServerError, err, "Can't delete tag")
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Deleted tag: %v", tag)})
}

// GenerateRandomQuestions returns the handler for GET /GenerateRandomQuestions
func (e *Env) GenerateRandomQuestions(c *gin.Context) {
	qs := d.Questions{
		Questions: []d.Question{
			{
				ID:       "1",
				Question: "q1",
				Answer:   "a1",
				Features: []string{"plain", "past", "conditional", "hallo", "good morning"},
			},
			{
				ID:       "2",
				Question: "q2",
				Answer:   "aa",
				Features: []string{"plain", "past", "q2"},
			},
			{
				ID:       "3",
				Question: "q3",
				Answer:   "a3",
				Features: []string{"polite", "present", "q3"},
			},
		},
	}

	c.JSON(http.StatusOK, qs)
}
