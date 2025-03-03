import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade vdoConverter = new VideoConversionFacade();
        String fileName = "lorem.ogg";
        String fileFormat = "mp4";

        File result = vdoConverter.convertVideo(fileName, fileFormat);

        System.out.println("conversion complete: " + result.getName());
    }
}

// 65070501076 Danai Saengbuamad