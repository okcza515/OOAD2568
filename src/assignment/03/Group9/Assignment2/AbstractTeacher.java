// Abstract base class for all school staff
abstract class AbstractTeacher implements SchoolStaffMember {
    private String name;

    public AbstractTeacher(String name) {
        this.name = name;
    }

    // Default implementations for common responsibilities
    @Override
    public void makeAnnouncements() {
        System.out.println("Making announcements to the class");
    }

    @Override
    public void takeAttendence() {
        System.out.println("Taking class attendance");
    }

    @Override
    public void collectPaperWork() {
        System.out.println("Collecting homework and assignments");
    }

    @Override
    public void conductHallwayDuties() {
        System.out.println("Monitoring hallways during designated times");
    }

    @Override
    public void performOtherResponsibilities() {
        System.out.println("Performing other school duties");
    }
}

// Regular teacher classes that can teach
class RegularTeacher extends AbstractTeacher implements TeachingCapable {
    private String subject;

    public RegularTeacher(String name, String subject) {
        super(name);
        this.subject = subject;
    }

    @Override
    public void teach() {
        System.out.println("Teaching " + subject);
    }
}
