package constants

// Constants for onboarding
const (
	BenifitOfWater    = "Вода жизненно важна для здоровья, так как поддерживает работу всех органов, выводит токсины и регулирует температуру тела. Она улучшает кожу, энергию и концентрацию, а также помогает пищеварению и снижает риск обезвоживания. Регулярное употребление воды – ключ к хорошему самочувствию."
	StandardNorm      = "Стандартная норма употребления воды для взрослого человека составляет около 2–2,5 литров в день. Однако точное количество зависит от возраста, пола, уровня физической активности и климатических условий."
	WaterNormQuestion = "Хочешь рассчитать свою или использовать стандартную норму?"
	WaterNormResult   = "Ваша норма употребления воды составляет %d миллилитров в день."
)

// Constants for onboarding first question
const (
	WaterNormSelfKey    = "self"
	WaterNormSelf       = "Рассчитать свою"
	WaterNormDefaultKey = "default"
	WaterNormDefault    = "Использовать стандартную"
)

// Constants for onboarding sex questrion
const (
	SexQuestion = "Выберите свой пол"
	SexKey      = "sex"
)

// Constants for onboarding physical activity
const (
	PhysicalActivityQuestion = "Выберите свою физическую активность"
	PhysicalActivityKey      = "physical_activity"
)

// Constants for onboarding climate
const (
	ClimateQuestion = "Выберите свой климат"
	ClimateKey      = "climate"
)

// Constants for onboarding weight
const (
	WeightQuestion        = "Введите свой вес в килограммах"
	WeightKey             = "weight"
	WeightErrorMessage    = "Пожалуйста, введите вес в килограммах в виде целого числа. Например: 75"
	WeightNotValidMessage = "Введено некорректное значение веса. Пожалуйста, повторите ввод (от 30 до 300)."
)

// Constants for onboarding water goal
const (
	WaterGoalKey     = "water_goal"
	WaterGoalDefault = 2000
)

// Constants for onboarding timezone
const (
	TimezoneQuestion = "Выберите свой часовой пояс"
	TimezoneKey      = "timezone"
)
