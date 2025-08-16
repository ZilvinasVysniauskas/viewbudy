// main.m
#import "bin/golib.h"
#import <Foundation/Foundation.h>

void playVideo(const char *path);

int main(int argc, const char *argv[]) {
  @autoreleasepool {
    StartGoLogic();

    if (argc < 2) {
      NSLog(@"Usage: %s /path/to/video.mp4", argv[0]);
      return 1;
    }

    const char *videoPath = argv[1];
    playVideo(videoPath);
  }
  return 0;
}
