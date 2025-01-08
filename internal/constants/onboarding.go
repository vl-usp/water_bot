package constants

// nolint:gochecknoglobals
const (
	BenifitOfWater    = "Вода жизненно важна для здоровья, так как поддерживает работу всех органов, выводит токсины и регулирует температуру тела. Она улучшает кожу, энергию и концентрацию, а также помогает пищеварению и снижает риск обезвоживания. Регулярное употребление воды – ключ к хорошему самочувствию."
	StandardNorm      = "Стандартная норма употребления воды для взрослого человека составляет около 2–2,5 литров в день. Однако точное количество зависит от возраста, пола, уровня физической активности и климатических условий."
	WaterNormQuestion = "Хочешь рассчитать свою или использовать стандартную норму?"
	WaterNormResult   = "Ваша норма употребления воды составляет %d миллилитров в день."
)

// nolint:gochecknoglobals
const (
	WaterNormSelfKey    = "self"
	WaterNormSelf       = "Рассчитать свою"
	WaterNormDefaultKey = "default"
	WaterNormDefault    = "Использовать стандартную"
)

// nolint:gochecknoglobals
const (
	SexQuestion  = "Выберите свой пол"
	SexKey       = "sex"
	SexMaleKey   = "male"
	SexMale      = "Мужской"
	SexFemaleKey = "female"
	SexFemale    = "Женский"
)

// nolint:gochecknoglobals
const (
	PhysicalActivityQuestion    = "Выберите свою физическую активность"
	PhysicalActivityKey         = "physical_activity"
	PhysicalActivityLowKey      = "low"
	PhysicalActivityLow         = "Низкая"
	PhysicalActivityModerateKey = "moderate"
	PhysicalActivityModerate    = "Умеренная"
	PhysicalActivityHighKey     = "high"
	PhysicalActivityHigh        = "Высокая"
)

// nolint:gochecknoglobals
const (
	ClimateQuestion     = "Выберите свой климат"
	ClimateKey          = "climate"
	ClimateColdKey      = "cold"
	ClimateCold         = "Холодный"
	ClimateTemperateKey = "temperate"
	ClimateTemperate    = "Умеренный"
	ClimateWarmKey      = "warm"
	ClimateWarm         = "Теплый"
	ClimateHotKey       = "hot"
	ClimateHot          = "Горячий"
)

// nolint:gochecknoglobals
const (
	WeightQuestion        = "Введите свой вес в килограммах"
	WeightKey             = "weight"
	WeightErrorMessage    = "Пожалуйста, введите вес в килограммах в виде целого числа. Например: 75"
	WeightNotValidMessage = "Введено некорректное значение веса. Пожалуйста, повторите ввод (от 30 до 300)."
)

// nolint:gochecknoglobals
const (
	WaterGoalKey     = "water_goal"
	WaterGoalDefault = 2000
)

// nolint:gochecknoglobals
const (
	TimezoneQuestion = "Выберите свой часовой пояс"
	TimezoneKey      = "timezone"
)
