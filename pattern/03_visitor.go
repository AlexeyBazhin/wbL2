package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/


/*
Определение
Поведенческий паттерн, который определяет операции, выполняемые над другими объектами.
Позволяет добавлять новые операции, не затрагивая классы, над которыми данные операции выполняются.

Простыми словами
Визитор отделяет операции от объектов, над которыми они выполняются.
Обходимся без даункастинга, в случае когда выносим общие методы, устанавливая аргументом общий интерфейс.

Применяется когда
- У объекта есть множество несвязанных операций
- Часто приходится добавлять новые операции
- Хочется вынести схожие операции разных объектов в одно место

Минусы
- Тесная связанность объектов (транзитивная)
- Применяют при наличии устоявшейся иерархии
*/
type (
	VisitorAccepter interface {
		AcceptVisitor(Visitor)
	}
	Visitor interface {
		VisitDog(Dog)
		VisitCat(Cat)
	}

	Dog struct{}
	Cat struct{}

	PetVisitor struct{}
	LogVisitor struct{}
)

func (dog Dog) AcceptVisitor(visitor Visitor) {
	visitor.VisitDog(dog)
}

func (cat Cat) AcceptVisitor(visitor Visitor) {
	visitor.VisitCat(cat)
}

func (visitor PetVisitor) VisitDog(dog Dog) {
	fmt.Printf("Глажу собаку %v\n", dog)
}

func (visitor PetVisitor) VisitCat(cat Cat) {
	fmt.Printf("Глажу кота %v\n", cat)
}

func (visitor LogVisitor) VisitDog(dog Dog) {
	log.Default().Println("Логгировал собаку")
}

func (visitor LogVisitor) VisitCat(cat Cat) {
	log.Default().Println("Логгировал Кота")
}

func main3() {
	dog := Dog{}
	cat := Cat{}
	dog.AcceptVisitor(PetVisitor{})
	cat.AcceptVisitor(LogVisitor{})
}