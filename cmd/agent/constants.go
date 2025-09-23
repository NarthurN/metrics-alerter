package main

const (
	// настройки

	// contentType по заданию
	contentType = "text/plain"

	// Типы метрик по заданию.
	gauge   = "gauge"
	counter = "counter"

	// метрики counter

	// PollCount — счётчик, увеличивающийся на 1 при каждом обновлении метрики из пакета runtime
	PollCount = "PollCount"

	// метрики gauge

	// RandomValue — обновляемое произвольное значение
	RandomValue = "RandomValue"

	// Alloc — объём байт, занятых «живыми» объектами в куче сейчас (текущая используемая память кучи).
	Alloc = "Alloc"

	// BuckHashSys — байты, занятые структурами хеш-таблиц профилировщика (buck hash) внутри рантайма.
	BuckHashSys = "BuckHashSys"

	// Frees — общее число операций освобождения памяти (количество free), накопительное.
	Frees = "Frees"

	// GCCPUFraction — доля CPU, потраченная сборщиком мусора с момента старта процесса (0.0..1.0).
	GCCPUFraction = "GCCPUFraction"

	// GCSys — байты, выделенные рантаймом под внутренние структуры GC.
	GCSys = "GCSys"

	// HeapAlloc — байты, реально занятые объектами в куче (как Alloc; ключевой индикатор «живых» данных).
	HeapAlloc = "HeapAlloc"

	// HeapIdle — байты кучи, не используемые приложением (idle), но ещё удерживаемые рантаймом у ОС.
	HeapIdle = "HeapIdle"

	// HeapInuse — байты кучи, находящиеся в использовании (in-use) под арены/страницы с объектами.
	HeapInuse = "HeapInuse"

	// HeapObjects — количество объектов в куче (живых) на момент чтения статистики.
	HeapObjects = "HeapObjects"

	// HeapReleased — байты, возвращённые из кучи обратно ОС (release), то есть реально освобождённые системе.
	HeapReleased = "HeapReleased"

	// HeapSys — суммарные байты виртуальной памяти, зарезервированные под кучу у ОС.
	HeapSys = "HeapSys"

	// LastGC — время (монотонные наносекунды от эпохи рантайма), когда завершился последний цикл GC.
	LastGC = "LastGC"

	// Lookups — количество поисков символов в таблицах профилировщика (stack/pprof) с момента старта.
	Lookups = "Lookups"

	// MCacheInuse — байты в использовании для per-P mcache (локальные кэши аллокатора).
	MCacheInuse = "MCacheInuse"

	// MCacheSys — байты, выделенные у ОС под структуры mcache.
	MCacheSys = "MCacheSys"

	// MSpanInuse — байты в использовании под структуры mspan (метаданные арен/страниц).
	MSpanInuse = "MSpanInuse"

	// MSpanSys — байты, выделенные у ОС под структуры mspan.
	MSpanSys = "MSpanSys"

	// Mallocs — общее число операций выделения памяти (количество malloc), накопительное.
	Mallocs = "Mallocs"

	// NextGC — целевой порог байт HeapAlloc, при достижении которого планируется следующий GC.
	NextGC = "NextGC"

	// NumForcedGC — число «форсированных» сборок мусора (вызванных вручную или по внутренним сигналам).
	NumForcedGC = "NumForcedGC"

	// NumGC — количество завершённых циклов GC с момента старта процесса.
	NumGC = "NumGC"

	// OtherSys — байты, занятые прочими внутренними структурами рантайма (не попавшие в иные категории).
	OtherSys = "OtherSys"

	// PauseTotalNs — суммарное время пауз GC в наносекундах, накопительно.
	PauseTotalNs = "PauseTotalNs"

	// StackInuse — байты стека горутин, находящиеся в использовании (in-use).
	StackInuse = "StackInuse"

	// StackSys — байты, зарезервированные у ОС под стеки горутин.
	StackSys = "StackSys"

	// Sys — общая память (байты), запрошенная рантаймом у ОС под все подсистемы (heap, stack, прочее).
	Sys = "Sys"

	// TotalAlloc — суммарный объём байт, когда-либо выделенных в куче с момента старта (не уменьшается).
	TotalAlloc = "TotalAlloc"
)
