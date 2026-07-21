---
title:  "Go Duck Yourself: Encapsulation in Ruby vs Go"
date:   2026-07-21 00:00:00 -0500
categories: web
layout: single
classes:
  - landing
  - dark-theme
---
![Encapsulation in Ruby vs Go](/assets/img/encapsulation-ruby-vs-go.jpg)
Some time ago I decided to give Go a go, being used to mostly dynamic/interpreted languages, mainly Ruby, I thought the transition was going to be hard but surprisingly, these 2 languages share a good amount of concepts. The most obvious one at least to me, was encapsulation.

# Classes vs Struct + Pointer Receivers

Probably the most common keyword used in Ruby to describe an Object that holds data and behavior is a class. If you are fan of Dave Thomas you may be also used to modules and structs but is fair to say that in Ruby classes have the lead:

```ruby
class Dog
	def initialize(name, breed)
		@breed = breed
		@name = name
	end

	def is_a_good_dog
		puts "#{@name} is the best dog ever, congrats!"
	end
end

dog = Dog.new("Noah", "Golden Retriever")
dog.is_a_good_dog
# => Noah is the best dog ever, congrats!
```

In Go we do not have classes but the intent of encapsulating behavior is present by leveraging Structs (data) and Pointer Receivers (behavior), so we can achieve a similar result, even though some knowledge on pointers is required (more on that later):

```go
package main

import "fmt"

type dog struct {
	name  string
	breed string
}

func newDog(name string, breed string) *dog {
	return &dog{name, breed}
}

func (dog *dog) isBestDog() {
	fmt.Printf("%s is the best dog, congrats!\n", dog.name)
}

func main() {
	var noah *dog = newDog("Noah", "Golden Retriever")
	noah.isBestDog()
}
// => Noah is the best dog, congrats!
```

Go also does not have a ‘constructor’, the newDog function above is common idiom in the language to handle initialization.

# Private accessibility

Assigning private accessibility to methods and attributes in Ruby is simple and beautiful:

```ruby
class Dog
	# initialize and more...
	
	private # or protected
	# private fields
	# private methods
end
```

This encapsulation in Ruby lets you choose what data and behavior is accessible from outside the class. Go does not have the same mechanism but a decent equivalent:

```go
package main

// Private structs and functions start with lowercase...
type dog struct {
	name  string
	breed string
}

// Public ones with uppercase
func NewDog(name string, breed string) *dog {
	return &dog{name, breed}
}
```

I do not think that Go’s solution is as good as Ruby’s since it probably means I need to read more code to find out whether something is accessible outside of a package but at least they have a solution that is simple, in theory.

# Global variables and Class methods

Doing globals in a Ruby class is quite simple too:

```ruby
class Dog
	BEST_DOG_BREEDS = ['Golden Retriever', 'Border Collie'].freeze
	def self.description
		puts "This is a class that describes a Dog."
	end
end
```

Very little code to achieve this behavior, Go is also very simple:

```go
package dog

import "fmt"

var BestDogBreeds = []string{"Golden Retriever", "Border Collie"}

func Description() {
	fmt.Println("This is a package that describes a Dog.")
} 
```

The code above allows us to call `Dog.description` or `Dog::BEST_DOG_BREEDS` in Ruby, and in Go we can import the `dog` package (`import “myproject/dog”`) and call `dog.BestDogBreeds` or `dog.Description()`. We do not need the `dog` namespace if we are on the same package though.

# Idioms: attr_accessor and pass by reference

And as expected, there are some little cool idioms present in both languages that are exclusive to them(I doubt these are the only languages with this though), in Ruby, we got `attr_accessor`:

```ruby
class Dog
	attr_accessor :name, :breed

	def initialize(name, breed)
		@name = name
		@breed = breed
	end
end

noel = Dog.new('john', 'Golden Retriever')
puts noel.name
# => 'john'
noel.name = 'Noel'
puts noel.name
# => 'Noel'
```

If you are not used to pointer referencing and dereferencing this might look like an annoying footgun, but Go lets you pass data to receivers and functions by either reference or copy. Which one is better? well, it depends but in Go you choose (Ruby defaults to passing by reference value):

```go
package dog

import "fmt"

type dog struct {
	name string
	breed string
}

func newDog(name string, breed string) *dog {
	return &dog{name, breed}
}

func (dog *dog) isBestDogByPointer() {
	// referenced by pointer
	fmt.Printf("%s is the best dog, congrats!\n", dog.name)
}

func (dog dog) isBestDogByValue() {
	// referenced by value i.e. shallow copy
	fmt.Printf("%s is the best dog, congrats!\n", dog.name)
}
```

Even though these 2 look different from the outside, after using them for a couple of programs there are some bridges that can be built between them; I will continue to explore those for sure.
Repo: https://github.com/christianpaez/go-duck-yourself
