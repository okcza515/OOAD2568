
public class Application {
	public static void main(String[] args) {
		String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";
		FileDataSource plain = new FileDataSource("OutputDemo.txt");
		plain.writeData(salaryRecords);
		System.out.println(plain.readData());
	}
}
