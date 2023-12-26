namespace hellmath {

// TODO: Task 1 - Define an `AccountStatus` enumeration to represent the four
// account types: `troll`, `guest`, `user`, and `mod`.
enum class AccountStatus {
  troll,
  guest,
  user,
  mod
};

// TODO: Task 1 - Define an `Action` enumeration to represent the three
// permission types: `read`, `write`, and `remove`.
enum class Action {
  read,
  write,
  remove
};

// TODO: Task 2 - Implement the `display_post` function, that gets two arguments
// of `AccountStatus` and returns a `bool`. The first argument is the status of
// the poster, the second one is the status of the viewer.
bool display_post (AccountStatus poster, AccountStatus viewer) {
  switch (poster) {
    case AccountStatus::troll:
      if (viewer == AccountStatus::troll) {
        return true;
      }
      return false;
    default:
      return true;
  }

  switch (viewer) {
    case AccountStatus::troll:
      if (poster == AccountStatus::troll) {
        return true;
      }
      return false;
    default:
      return true;
  }
}
// TODO: Task 3 - Implement the `permission_check` function, that takes an
// `Action` as a first argument and an `AccountStatus` to check against. It
// should return a `bool`.
bool permission_check(Action action, AccountStatus status) {
  switch (status) {
    case AccountStatus::guest:
      switch(action) {
        case Action::read:
          return true;
        default:
          return false;
      }
    case AccountStatus::troll:
    case AccountStatus::user:
      switch (action)
      {
      case Action::read:
        return true;
      case Action::write:
        return true;
      default:
        return false;
      }
    default:
      return true;
  }
}
// TODO: Task 4 - Implement the `valid_player_combination` function that
// checks if two players can join the same game. The function has two parameters
// of type `AccountStatus` and returns a `bool`.
bool valid_player_combination(AccountStatus a, AccountStatus b) {
  if (a == AccountStatus::guest || b == AccountStatus::guest) {
    return false;
  }

  if (a == AccountStatus::troll && b == AccountStatus::troll) {
    return true;
  }

  if ((a == AccountStatus::troll && b != AccountStatus::troll) || (a != AccountStatus::troll && b == AccountStatus::troll)) {
    return false;
  }
  return true;
}
// TODO: Task 5 - Implement the `has_priority` function that takes two
// `AccountStatus` arguments and returns `true`, if and only if the first
// account has a strictly higher priority than the second.
bool has_priority(AccountStatus a, AccountStatus b) {
  switch (a)
  {
  case AccountStatus::mod:
    if (b != AccountStatus::mod) {
      return true;
    }
    return false;
  case AccountStatus::user:
    switch(b) {
      case AccountStatus::mod:
        return false;
      case AccountStatus::user:
        return false;
      default:
        return true;
    }
  case AccountStatus::guest:
    switch (b)
    {
    case AccountStatus::troll:
      return true;    
    default:
      return false;
    }
  default:
    return false;
  }
}

}  // namespace hellmath