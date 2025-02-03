// Base class for all teachers
public abstract class Teacher {
    protected String name;

    public Teacher(String name) {
        this.name = name;
    }

    public void makeAnnouncements() {
        System.out.println(name + " is making announcements.");
    }

    public void takeAttendance() {
        System.out.println(name + " is taking attendance.");
    }

    public void collectPaperWork() {
        System.out.println(name + " is collecting paperwork.");
    }

    public void conductHallwayDuties() {
        System.out.println(name + " is conducting hallway duties.");
    }

    public void performOtherResponsibilities() {
        System.out.println(name + " is performing other responsibilities.");
    }
}
