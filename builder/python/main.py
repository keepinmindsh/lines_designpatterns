from __future__ import annotations
from abc import ABC, abstractmethod, ABCMeta
from typing import Any


class Builder(ABC):

    @property
    @abstractmethod
    def dropship(self) -> None:
        pass

    @abstractmethod
    def get_on_the_ship(self) -> None:
        pass

    @abstractmethod
    def get_off_the_ship(self) -> None:
        pass


class ConcreteBuilder1(Builder):

    def __init__(self) -> None:
        self.reset()

    def reset(self) -> None:
        self._product = Product1()

    @property
    def dropship(self) -> Product1:
        product = self._product
        self.reset()
        return product

    def get_on_the_ship(self) -> None:
        self._product.add("PartA1")

    def get_off_the_ship(self) -> None:
        self._product.add("PartB1")


class Product1:
    def __init__(self) -> None:
        self.parts = []

    def add(self, part: Any) -> None:
        self.parts.append(part)

    def list_parts(self) -> None:
        print(f"Product parts: {', '.join(self.parts)}", end="")


class Director:
    def __init__(self) -> None:
        self._builder = None

    @property
    def builder(self) -> Builder:
        return self._builder

    @builder.setter
    def builder(self, builder: Builder) -> None:
        self._builder = builder

    def build_minimal_viable_product(self) -> None:
        self.builder.get_on_the_ship()

    def build_full_featured_product(self) -> None:
        self.builder.get_on_the_ship()
        self.builder.get_on_the_ship()
        self.builder.get_off_the_ship()


if __name__ == "__main__":

    director = Director()
    builder = ConcreteBuilder1()
    director.builder = builder

    print(type(director))

    print("Standard basic product: ")
    director.build_minimal_viable_product()
    builder.dropship.list_parts()

    print("\n")

    print("Standard full featured product: ")
    director.build_full_featured_product()
    builder.dropship.list_parts()

    print("\n")

    # Remember, the Builder pattern can be used without a Director class.
    print("Custom product: ")
    builder.get_on_the_ship()
    builder.get_off_the_ship()
    builder.dropship.list_parts()