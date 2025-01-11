package gotestify

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Общее количество кафе
	totalCount := 4

	// Создаем запрос с count больше, чем доступно кафе
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	// Создаем фиктивный записывающий ответ (ResponseRecorder)
	responseRecorder := httptest.NewRecorder()

	// Указываем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа - 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	// Получаем тело ответа
	body := responseRecorder.Body.String()

	// Проверяем, что длина возвращенного списка кафе равна общему количеству кафе
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount, "Response should contain all available cafes")

	// Проверяем, что список кафе соответствует ожидаемому
	expectedCafes := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	assert.Equal(t, expectedCafes, list, "Returned cafe list does not match expected")
}
