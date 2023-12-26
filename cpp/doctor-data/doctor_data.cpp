
#include "doctor_data.h"
#include<string>

namespace heaven {
  Vessel::Vessel(std::string name, int generation) {
    this->name = name;
    this->generation = generation;
  }
  Vessel::Vessel(std::string name, int generation, star_map::System current_system) {
    this->name = name;
    this->generation = generation;
    this->current_system = current_system;
  }
  Vessel Vessel::replicate(std::string name) {
    Vessel newvessel{name, this->generation + 1, this->current_system};
    return newvessel;
  }

  void Vessel::make_buster() {
    this->busters += 1;
  }

  bool Vessel::shoot_buster() {
    if (this->busters > 0) {
      this->busters = 0;
      return true;
    } else {
      this->busters = 0;
      return false;
    }
  }

  std::string get_older_bob(Vessel bob1, Vessel bob2) {
    if (bob1.generation > bob2.generation) {
      return bob2.name;
    } else {
      return bob1.name;
    }
  }
  bool in_the_same_system(Vessel bob1, Vessel bob2) {
    return bob1.current_system == bob2.current_system;
  }
}

