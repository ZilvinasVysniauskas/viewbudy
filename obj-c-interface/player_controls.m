#include "gobridge.h"
#import <AVFoundation/AVFoundation.h>
#import <AVKit/AVKit.h>
#import <AppKit/AppKit.h>

static AVPlayer *gPlayer = nil;

void playVideo(const char *path) {
  @autoreleasepool {

    // Log the start of the function and the path received.
    NSString *msg =
        [NSString stringWithFormat:@"playVideo called with path: %s", path];

    // Convert NSString to UTF-8 C string and pass to logMessage
    logMessage((char *)[msg UTF8String]);

    NSString *nsPath = [NSString stringWithUTF8String:path];
    NSURL *url = [NSURL fileURLWithPath:nsPath];

    gPlayer = [AVPlayer playerWithURL:url];
    AVPlayerView *playerView =
        [[AVPlayerView alloc] initWithFrame:NSMakeRect(0, 0, 800, 600)];
    [playerView setPlayer:gPlayer];

    NSApplication *app = [NSApplication sharedApplication];
    [app setActivationPolicy:NSApplicationActivationPolicyRegular];

    NSWindow *window =
        [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, 800, 600)
                                    styleMask:(NSWindowStyleMaskTitled |
                                               NSWindowStyleMaskClosable |
                                               NSWindowStyleMaskResizable)
                                      backing:NSBackingStoreBuffered
                                        defer:NO];
    [window setContentView:playerView];
    [window makeKeyAndOrderFront:nil];

    [app activateIgnoringOtherApps:YES];
    [gPlayer play];

    // The app run loop is a blocking call, so this log won't show up until it's
    // finished.
    [app run];
  }
}

void pauseVideo(void) {
  @autoreleasepool {
    if (gPlayer) {
      [gPlayer pause];
      logMessage("Video playback paused.");
    } else {
      logMessage("Cannot pause: no player instance available.");
    }
  }
}

void processEvents(void) {
    @autoreleasepool {
        NSApplication *app = [NSApplication sharedApplication];
        NSEvent *event;

        // Process one batch of events, non-blocking
        while ((event = [app nextEventMatchingMask:NSEventMaskAny
                                         untilDate:[NSDate dateWithTimeIntervalSinceNow:0.01]
                                            inMode:NSDefaultRunLoopMode
                                           dequeue:YES])) {
            [app sendEvent:event];
        }
        [app updateWindows];
    }
}
