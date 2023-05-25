# ascii.live

A project for hosting curl-based animations, all in one place, and a follow up to [`parrot.live`](https://github.com/hugomd/parrot.live).

Any animations you want to add are welcome! 🎉

Try it out in your terminal:

```bash
curl https://ascii-bash.onrender.com/parrot
```

## Running locally
To run the server locally on port `8080`, run:
```bash
go run main.go
```

## Adding frames
* Fork this repository
* Create a new frame file in [`/frames`](./frames), call it the name of your frames/animation, e.g. `parrot.go`
* Inside your new file, create an exported list of frames, e.g.

```Golang
package frames

// This is the value stored in the FrameMap
var MyAnimation = DefaultFrameType(myAnimationFrames)

var myAnimationFrames = []string{
  `Frame1`,
  `Frame2`,
  `Frame3`,
}
```

* In [`./frames/frames.go`](./frames/frames.go), add your animation to the `FrameMap`
* Commit and push your changes, and make a PR! If this is your first time making a PR, [check GitHub's help page on the topic](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)