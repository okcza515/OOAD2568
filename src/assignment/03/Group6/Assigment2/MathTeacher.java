
public class MathTeacher extends Teacher implements Teachable{

	@Override
	public void teach() {
		System.out.println("Taught Math");
	}

	public void takeAttendence() {
		System.out.println("Math teacher took attendence..");
	}


}
