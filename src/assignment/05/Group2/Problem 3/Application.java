public class Application {
	public static void main(String[] args) {
		String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";

		// Plain data
		Datasource source = new FileDataSource("OutputDemo.txt");
		source.writeData(salaryRecords);
		System.out.println("Plain data:");
		System.out.println(source.readData());

		// Encrypted data
		Datasource encrypted = new EncryptionDecorator(
				new FileDataSource("OutputDemo_encrypted.txt"));
		encrypted.writeData(salaryRecords);
		System.out.println("\nEncrypted data:");
		System.out.println(encrypted.readData());

		// Compressed data
		Datasource compressed = new CompressionDecorator(
				new FileDataSource("OutputDemo_compressed.txt"));
		compressed.writeData(salaryRecords);
		System.out.println("\nCompressed data:");
		System.out.println(compressed.readData());

		// Both encryption and compression
		Datasource encryptedAndCompressed = new EncryptionDecorator(
				new CompressionDecorator(
						new FileDataSource("OutputDemo_enc_comp.txt")));
		encryptedAndCompressed.writeData(salaryRecords);
		System.out.println("\nEncrypted and compressed data:");
		System.out.println(encryptedAndCompressed.readData());
	}
}
// Ratchanon Tarawan 65070503464