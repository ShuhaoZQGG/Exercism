class Badge {
    public String print(Integer id, String name, String department) {
        String badge = id == null ? "" : "[" + id + "]";
        department = department == null ? "OWNER" : department.toUpperCase();

        if (badge != "") {
            return badge + " - " + name + " - " + department;
        }

        return name + " - " + department;
    }
}
