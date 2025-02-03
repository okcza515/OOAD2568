// Math Teacher
public class MathTeacher extends Teacher implements Teachable {
    public MathTeacher(String name) {
        super(name);
    }

    @Override
    public void teach() {
        System.out.println(name + " is teaching Math.");
    }
}

