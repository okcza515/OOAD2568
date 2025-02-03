// SubstituteTeacher does NOT implement Teachable
public class SubstituteTeacher extends Teacher {
    public SubstituteTeacher(String name) {
        super(name);
    }

    // No teach() method, follows Liskov principle
    public void assist() {
        System.out.println(name + " is assisting but not teaching.");
    }
}
