package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(cond bool) validation.RuleFunc {
	// validation.RuleFunc — это функция-валидатор, которая будет вызвана для проверки поля
	// Внутри возвращённой функции:
	// Если cond == true → проверяет, что value не пусто (вызов validation.Required)
	// Если cond == false → просто возвращает nil (ошибки нет, валидация пройдена)
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
