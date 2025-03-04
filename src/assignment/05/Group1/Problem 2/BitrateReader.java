public class BitrateReader {
    public static VideoFile read(VideoFile file, Codec codec) {
        System.out.println("BitrateReader: reading file...");
        return file;
    }

    public static VideoFile convert(VideoFile buffer, Codec codec) {
        System.out.println("BitrateReader: writing file...");
        VideoFile converted_file = new VideoFile(buffer.getNameWithOutCodec() + "." + codec.getType());
        return converted_file;
    }
}