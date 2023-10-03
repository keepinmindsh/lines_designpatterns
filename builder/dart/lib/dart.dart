enum UnitType { Marine, Medic, FireBat }

class Builder {
  String name;
  int age;
  UnitType type;

  Builder({required this.name, required this.age, required this.type});
}

class Barrack {
  makeUnit(Builder builder){
    switch(builder.type){
      case UnitType.Marine :
        return Marine();
      case UnitType.FireBat :
        return FireBat();
      case UnitType.Medic :
        return Medic();
    }
  }
}

abstract interface class Unit {
  void attack();
  void defence();
}

class Marine implements Unit {
  @override
  void attack() {
    print("Marine Attack!");
  }

  @override
  void defence() {
    print("Marine Defence!");
  }
}

class Medic implements Unit {
  @override
  void attack() {
    // TODO: implement attack
  }

  @override
  void defence() {
    // TODO: implement defence
  }
}

class FireBat implements Unit {
  @override
  void attack() {
    // TODO: implement attack
  }

  @override
  void defence() {
    // TODO: implement defence
  }

}

void main(){
  var builder = Builder(name: "Lines", age: 500, type: UnitType.Marine);

  var barrack = Barrack();

  Unit makeUnit = barrack.makeUnit(builder);

  makeUnit.attack();
  makeUnit.defence();
}
