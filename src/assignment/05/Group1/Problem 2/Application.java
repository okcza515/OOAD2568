import java.io.File;
public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        
        // ใส่ชื่อไฟล์ตรงนี้
        String inputFileName = "group_one.ogg";

        VideoFile file = new VideoFile(inputFileName);

        // ดึงนามสกุลไฟล์เดิมออกมา จากนั้นส่งไปให้ converter แปลงเป็นอีกนามสกุล
        File convertedFile = converter.convertVideo(inputFileName, file.getCodecType());

        if (convertedFile != null) {
            System.out.println("Conversion process completed.");
            System.out.println("-------------------------------------------------------------------------------------");
            System.out.println("Output file: " + convertedFile.getName()); // แสดงชื่อไฟล์ที่แปลงเสร็จแล้ว
        } else {
            System.out.println("Conversion process failed.");
        }
    }
}