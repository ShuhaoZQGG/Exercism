abstract class Fighter {
    public String toString() {
        return String.format("Fighter is a %s", this.getClass().getSimpleName());
    }
    boolean isVulnerable() {
        return false;
    }

    abstract int damagePoints(Fighter fighter);

}

class Warrior extends Fighter {

    @Override
    public String toString() {
        return "Fighter is a Warrior";
    }

    @Override
    int damagePoints(Fighter fighter) {
        return fighter.isVulnerable() ? 10 : 6;
    }
}

class Wizard extends Fighter {
    boolean spellPrepared = false;
    @Override
    boolean isVulnerable() {
        return !this.spellPrepared;
    }

    @Override
    int damagePoints(Fighter fighter) {
        return this.spellPrepared ? 12 : 3;
    }

    void prepareSpell() {
        this.spellPrepared = true;
    }

}
