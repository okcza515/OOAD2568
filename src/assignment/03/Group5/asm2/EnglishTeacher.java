
// English Teacher
public class EnglishTeacher extends Teacher implements Teachable {
    public EnglishTeacher(String name) {
        super(name);
    }

    @Override
    public void teach() {
        System.out.println(name + " is teaching English.");
    }
}