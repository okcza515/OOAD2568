public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();

        // Convert OGG to MP4
        System.out.println("Application: Converting OGG to MP4...");
        converter.convertVideo("video.ogg", "mp4");

        // Convert MP4 to OGG
        System.out.println("Application: Converting MP4 to OGG...");
        converter.convertVideo("video.mp4", "ogg");

        System.out.println("Application: Conversion process finished.");
    }
}
