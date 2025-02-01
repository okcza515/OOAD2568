public interface Iteach {
    void teach();
}

public abstract class Teacher {
    private String name;
    
    public void makeAnnouncements(){
        System.out.println("Making announcement");
    }
    public void takeAttendance(){
        System.out.println("Taking attendance");
    }
    public void collectPaperWork(){
        System.out.println("Collecting paperwork");
    }
    public void conductHallwayDuties(){
        System.out.println("Conducting hallway duty");
    }
    public void performOtherResponsibilities(){
        System.out.println("Performing other responsibilities");
    }
}

public class MathTeacher extends Teacher implements Iteach {
    public void teach(){
        System.out.println("Teaching Math");
    }
}

public class EnglishTeacher extends Teacher implements Iteach {
    public void teach(){
        System.out.println("Teaching English");
    }
}

public class ScienceTeacher extends Teacher implements Iteach {
    public void teach(){
        System.out.println("Teaching Science");
    }
}

public class SubstituteTeacher extends Teacher {
    
}
//65070501069