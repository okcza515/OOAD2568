
public class EnglishTeacher extends Teacher implements Teachable {

	@Override
	public void teach() {
		System.out.println("Taught English");
	}

	public void makeAnnouncements() {
		System.out.println("Eng teacher made annoucement");
	}

}
