import 'package:starcraft/starcraft.dart' as starcraft;

void main(List<String> arguments) {
  print('Hello world: ${starcraft.calculate()}!');
}

class Car {
  Car(this.engine);

  // field
  String engine = "E1001";

  // function
  void disp() {
    print(engine);
  }
}

class Student {
  String name;
  int age;

  String get stud_name {
    return name;
  }

  void set stud_name(String name) {
    this.name = name;
  }

  void set stud_age(int age) {
    if ( age <= 0 ) {
      print("Age should be greader than 5");
    }else {
      this.age = age;
    }

  }

  int get stud_age {
    return age;
  }
}

