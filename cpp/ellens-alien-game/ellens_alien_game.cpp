namespace targets {
// TODO: Insert the code for the alien class here
  class Alien {
    public:
      int x_coordinate {};
      int y_coordinate {};

      Alien(int _x_coordinate, int _y_coordinate) {
        x_coordinate = _x_coordinate;
        y_coordinate = _y_coordinate;
      }

      int get_health() {
        return health;
      }

      bool is_alive() {
        if (health > 0) {
          return true;
        }
        return false;
      }

      bool hit() {
        if (is_alive()) {
          health -= 1;
          return true;
        }
        return false;
      }

      bool teleport(int new_x, int new_y) {
        x_coordinate = new_x;
        y_coordinate = new_y;
        return true;
      }

      bool collision_detection(Alien anotherAlien) {
        if (x_coordinate == anotherAlien.x_coordinate && y_coordinate == anotherAlien.y_coordinate) {
          return true;
        }
        return false;
      }
    private:
      int health {3};
  };
}  // namespace targets