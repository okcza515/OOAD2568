//Warapol Pratumta 65070503466
interface SchoolStaffMember {
    void makeAnnouncements();
    void takeAttendence();
    void collectPaperWork();
    void conductHallwayDuties();
    void performOtherResponsibilities();
}

// Interface for teaching capability
interface TeachingCapable {
    void teach();
}

// Sawitt Ngamvilaisiriwong 65070503469
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

// Chayaphon Chaisangkha 65070503409
// Specific subject teachers
class MathTeacher extends RegularTeacher {
    public MathTeacher(String name) {
        super(name, "Mathematics");
    }
}

class EnglishTeacher extends RegularTeacher {
    public EnglishTeacher(String name) {
        super(name, "English");
    }
}

class ScienceTeacher extends RegularTeacher {
    public ScienceTeacher(String name) {
        super(name, "Science");
    }
}
//Paratthakon Suksukhon
// Substitute teacher that doesn't implement TeachingCapable
class SubstituteTeacher extends AbstractTeacher {
    public SubstituteTeacher(String name) {
        super(name);
    }

    // Substitute specific methods
    public void followLessonPlan() {
        System.out.println("Following the lesson plan left by the regular teacher");
    }
}