
public class Application {
	public static void main(String[] args) {
		String salaryRecords = "Name,Salary\nThanopol Samsan,100000\nSteven Jobs,210000";

		DataSource source = new FileDataSource("OutputDemo.txt"); 
		source.writeData(salaryRecords);
		System.out.println("Normal data:");
		System.out.println(source.readData());

		DataSource encrypted = new EncryptionDecorator(
				new FileDataSource("OutputDemo_encrypted.txt")
		);
		encrypted.writeData(salaryRecords);
		System.out.println("\nEncrypted data:");
		System.out.println(encrypted.readData());

		DataSource compressed = new CompressionDecorator(
				new FileDataSource("OutputDemo_compressed.txt")
		);
		System.out.println("\nCompressed data:");
		compressed.writeData(salaryRecords);
		System.out.println(compressed.readData());

		DataSource encryptedAndCompressed = new EncryptionDecorator(
				new CompressionDecorator(
						new FileDataSource("OutputDemo_enc_comp.txt")
				)
		);
		System.out.println("\nEncrypted and compressed data:");
		encryptedAndCompressed.writeData(salaryRecords);
		System.out.println(encryptedAndCompressed.readData());
	}
}

//65070501088 Sopida Keawjongkool
