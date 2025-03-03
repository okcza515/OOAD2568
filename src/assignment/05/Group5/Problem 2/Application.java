//Natchanon 65070501018
class Application {
    public File convert(String filename, String format) {
        File file = new VideoFile(filename);
        String sourceCodec = CodecFactory.extract(file);

        Codec destinationCodec;
        if (format == "mp4") {
            destinationCodec = new MPEG4CompressionCodec();
        } else {
            destinationCodec = new OggCompressionCodec();
        }

        String buffer = BitrateReader.read(filename, sourceCodec);
        String result = BitrateReader.convert(buffer, destinationCodec);
        result = new AudioMixer().fix(result);

        return new File(result);
    }
}



