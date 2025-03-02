
public class Application {
	public static void main(String[] args) {
		String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";
		
		FileDataSource plain = new FileDataSource("OutputDemo.txt");
		plain.writeData(salaryRecords);
		System.out.println(plain.readData());

		FileDataSource encrypted = new EncryptionDecorator(new FileDataSource("OutputEncrypted.txt"));
        encrypted.writeData(salaryRecords);
        System.out.println(encrypted.readData());

		FileDataSource compressed = new CompressionDecorator(new FileDataSource("OutputCompressed.txt"));
        compressed.writeData(salaryRecords);
        System.out.println(compressed.readData());

		FileDataSource encryptedCompressed = new EncryptionDecorator(new CompressionDecorator(new FileDataSource("OutputCombined.txt")));
        encryptedCompressed.writeData(salaryRecords);
        System.out.println(encryptedCompressed.readData());
	}
}
