package textualContent

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"path/filepath"

	"github.com/hassannmoussaa/pill.go/clean"
)

type TextualContent map[string]map[string]string

var textualContentDirPath string
var textualContent *TextualContent

func Init(txtDirPath string) {
	if textualContent == nil {
		textualContentDirPath = filepath.FromSlash(txtDirPath)
		textualContent = &TextualContent{}
		//parse config file
		txtFile, err := os.Open(filepath.Join(textualContentDirPath, "txt.json"))
		defer txtFile.Close()
		if err == nil {
			txtFileInfo, _ := txtFile.Stat()
			size := txtFileInfo.Size()
			data := make([]byte, size)
			reader := bufio.NewReader(txtFile)
			readLen, _ := reader.Read(data)
			data = data[:readLen]
			err = json.Unmarshal(data, textualContent)
			if err != nil {
				clean.Error(err)
				os.Exit(1)
			}
		} else {
			clean.Error(err)
			os.Exit(1)
		}
	}
}

func Get() *TextualContent {
	if textualContent == nil {
		Init(textualContentDirPath)
	}
	return textualContent
}

func (this *TextualContent) GetGroup(key string) map[string]string {
	if key != "" {
		key = strings.TrimSpace(key)
		for k, v := range *this {
			if k == key {
				return v
			}
		}
	}
	return nil
}

func (this *TextualContent) GetOne(key1 string, key2 string) string {
	if key1 != "" && key2 != "" {
		key1 = strings.TrimSpace(key1)
		key2 = strings.TrimSpace(key2)
		txtGroup := this.GetGroup(key1)
		if txtGroup != nil {
			for k, v := range txtGroup {
				if k == key2 {
					return v
				}
			}
		}
	}
	return ""
}

func (this *TextualContent) OfFieldPlaceholder(key string) string {
	return this.GetOne("fields_placeholders", key)
}
func (this *TextualContent) OfEmptyContentPhrase(key string) string {
	return this.GetOne("empty_content_phrases", key)
}
func (this *TextualContent) OfFieldLabel(key string) string {
	return this.GetOne("fields_labels", key)
}
func (this *TextualContent) OfTitle(key string) string {
	return this.GetOne("titles", key)
}
func (this *TextualContent) OfTab(key string) string {
	return this.GetOne("tabs", key)
}
func (this *TextualContent) OfMenuItem(key string) string {
	return this.GetOne("menus_items", key)
}
func (this *TextualContent) OfTH(key string) string {
	return this.GetOne("th", key)
}
func (this *TextualContent) OfButton(key string) string {
	return this.GetOne("buttons", key)
}
func (this *TextualContent) OfAlbumType(key string) string {
	return this.GetOne("album_types", key)
}
func (this *TextualContent) OfSuccessfulMsg(key string) string {
	return this.GetOne("successful_messages", key)
}
func (this *TextualContent) OfFailureMsg(key string) string {
	return this.GetOne("failure_messages", key)
}
func (this *TextualContent) OfErrorMsg(key string) string {
	msg := this.GetOne("error_messages", key)
	if msg == "" {
		this.GetOne("error_messages", "operation_failed")
	}
	return msg
}

func (this *TextualContent) OfValidationMsg(key string) string {
	return this.GetOne("validation_messages", key)
}

func (this *TextualContent) OfMeta(key string) string {
	return this.GetOne("metas", key)
}

func (this *TextualContent) OfAdminPermission(key string) string {
	return this.GetOne("admin_permissions", key)
}

func GetGroup(key string) map[string]string {
	if textualContent != nil {
		return textualContent.GetGroup(key)
	}
	return nil
}

func GetOne(key1 string, key2 string) string {
	if textualContent != nil {
		return textualContent.GetOne(key1, key2)
	}
	return ""
}

func OfFieldPlaceholder(key string) string {
	if textualContent != nil {
		return textualContent.OfFieldPlaceholder(key)
	}
	return ""
}

func OfEmptyContentPhrase(key string) string {
	if textualContent != nil {
		return textualContent.OfEmptyContentPhrase(key)
	}
	return ""
}
func OfFieldLabel(key string) string {
	if textualContent != nil {
		return textualContent.OfFieldLabel(key)
	}
	return ""
}
func OfTitle(key string) string {
	if textualContent != nil {
		return textualContent.OfTitle(key)
	}
	return ""
}
func OfTab(key string) string {
	if textualContent != nil {
		return textualContent.OfTab(key)
	}
	return ""
}
func OfMenuItem(key string) string {
	if textualContent != nil {
		return textualContent.OfMenuItem(key)
	}
	return ""
}

func OfTH(key string) string {
	if textualContent != nil {
		return textualContent.OfTH(key)
	}
	return ""
}
func OfButton(key string) string {
	if textualContent != nil {
		return textualContent.OfButton(key)
	}
	return ""
}
func OfAlbumType(key string) string {
	if textualContent != nil {
		return textualContent.OfAlbumType(key)
	}
	return ""
}
func OfAdminPermission(key string) string {
	if textualContent != nil {
		return textualContent.OfAdminPermission(key)
	}
	return ""
}
func OfSuccessfulMsg(key string) string {
	if textualContent != nil {
		return textualContent.OfSuccessfulMsg(key)
	}
	return ""
}
func OfFailureMsg(key string) string {
	if textualContent != nil {
		return textualContent.OfFailureMsg(key)
	}
	return ""
}
func OfErrorMsg(key string) string {
	if textualContent != nil {
		return textualContent.OfErrorMsg(key)
	}
	return ""
}
func OfValidationMsg(key string) string {
	if textualContent != nil {
		return textualContent.OfValidationMsg(key)
	}
	return ""
}
func OfMeta(key string) string {
	if textualContent != nil {
		return textualContent.OfMeta(key)
	}
	return ""
}
