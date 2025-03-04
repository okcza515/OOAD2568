
public class Application {
	public static void main(String[] args) {
		String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";
		FileData plain = new FileDataSource("OutputDemo.txt");
		plain.writeData(salaryRecords);
		System.out.println(plain.readData());

		System.out.println("----------------------------");
		// To encrypt the data
		EncryptionDecorator encrypted = new EncryptionDecorator(plain);
		encrypted.writeData(salaryRecords);
		FileData encryptedFile = new FileDataSource("OutputDemo.txt");
		System.out.println(encryptedFile.readData());

		System.out.println("----------------------------");
		// To compress the data
		CompressionDecorator compressed = new CompressionDecorator(plain);
		compressed.writeData(salaryRecords);
		FileData compressedFile = new FileDataSource("OutputDemo.txt");
		System.out.println(compressedFile.readData());
	}
}
