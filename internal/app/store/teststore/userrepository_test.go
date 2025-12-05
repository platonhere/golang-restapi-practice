package teststore_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStore создаёт тестовую базу данных и проверяет успешное открытие соединения.
// Возвращает объект Store и функцию teardown, которая очищает указанные таблицы
// и связанные с н ими данные после теста.
// В тесте вызывается метод Create репозитория UserRepository и проверяется,
// что ошибка отсутствует, а возвращённый пользователь не nil.
func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// TestStore открывает соединение с тестовой БД, defer teardown("users") очищает таблицу после теста.
// Тест проверяет метод FindByEmail репозитория UserRepository.
// Сначала проверяется, что поиск пользователя в пустой таблице возвращает ошибку.
// Затем создаётся пользователь с указанным email, и повторный поиск должен пройти успешно:
// ошибки нет, а возвращённый пользователь не nil.
func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	email = u.Email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)

}
