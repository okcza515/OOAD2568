import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        File file = converter.convertVideo(".DS_Store", "mp4");
        System.out.println("\n" + file);
    }
}

// 65070501085