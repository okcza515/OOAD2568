
public class ScienceTeacher extends Teacher implements Teachable {

	@Override
	public void teach() {
		System.out.println("taught science");
	}

	public void collectPaperWork() {
		System.out.println("Eng teacher collected paperwork..");
	}

}
