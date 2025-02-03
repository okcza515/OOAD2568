// Science Teacher
public class ScienceTeacher extends Teacher implements Teachable {
    public ScienceTeacher(String name) {
        super(name);
    }

    @Override
    public void teach() {
        System.out.println(name + " is teaching Science.");
    }
}