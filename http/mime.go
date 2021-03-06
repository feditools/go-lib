package http

type (
	Mime   string
	Suffix string
)

func (m Mime) String() string {
	return string(m)
}

func (s Suffix) String() string {
	return string(s)
}

const (
	// MimeAll matches any mime type.
	MimeAll Mime = `*/*`

	// MimeAppJRDJSON represents a JSON Resource Descriptor type.
	MimeAppJRDJSON Mime = `application/jrd+json`
	// MimeAppJSON represents a JavaScript Object Notation type.
	MimeAppJSON Mime = `application/json`
	// MimeAppActivityJSON represents a JSON activity pub action type.
	MimeAppActivityJSON Mime = `application/activity+json`
	// MimeAppActivityLDJSON represents JSON-based Linked Data for activity streams type.
	MimeAppActivityLDJSON Mime = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
	// MimeImageGIF represents a gif image type.
	MimeImageGIF Mime = `image/gif`
	// MimeImageJPG represents a jpg image type.
	MimeImageJPG Mime = `image/jpeg`
	// MimeImagePNG represents a png image type.
	MimeImagePNG Mime = `image/png`
	// MimeImageSVG represents a svg image type.
	MimeImageSVG Mime = `image/svg+xml`
	// MimeImageWebP represents a webp image type.
	MimeImageWebP Mime = `image/webp`
	// MimeTextCSV represents a csv type.
	MimeTextCSV Mime = `text/csv`
	// MimeTextHTML represents a html type.
	MimeTextHTML Mime = `text/html`
	// MimeTextPlain represents a plain text type.
	MimeTextPlain Mime = `text/plain`

	// SuffixAppJSON represents a JavaScript Object Notation suffix.
	SuffixAppJSON Suffix = `json`
	// SuffixImageGIF represents a gif image suffix.
	SuffixImageGIF Suffix = `gif`
	// SuffixImageJPG represents a jpg image suffix.
	SuffixImageJPG Suffix = `jpg`
	// SuffixImagePNG represents a png image suffix.
	SuffixImagePNG Suffix = `png`
	// SuffixImageSVG represents a svg image suffix.
	SuffixImageSVG Suffix = `svg`
	// SuffixImageWebP represents a webp image suffix.
	SuffixImageWebP Suffix = `webp`
	// SuffixTextCSV represents a csv suffix.
	SuffixTextCSV Suffix = `csv`
	// SuffixTextHTML represents a html suffix.
	SuffixTextHTML Suffix = `html`
	// SuffixTextPlain represents a plain text suffix.
	SuffixTextPlain Suffix = `txt`
)

var (
	suffixToMime = map[Suffix]Mime{
		SuffixAppJSON:   MimeAppJSON,
		SuffixImageGIF:  MimeImageGIF,
		SuffixImageJPG:  MimeImageJPG,
		SuffixImagePNG:  MimeImagePNG,
		SuffixImageSVG:  MimeImageSVG,
		SuffixImageWebP: MimeImageWebP,
		SuffixTextCSV:   MimeTextCSV,
		SuffixTextHTML:  MimeTextHTML,
		SuffixTextPlain: MimeTextPlain,
	}

	mimeToSuffix = map[Mime]Suffix{
		MimeAppJSON:   SuffixAppJSON,
		MimeImageGIF:  SuffixImageGIF,
		MimeImageJPG:  SuffixImageJPG,
		MimeImagePNG:  SuffixImagePNG,
		MimeImageSVG:  SuffixImageSVG,
		MimeImageWebP: SuffixImageWebP,
		MimeTextCSV:   SuffixTextCSV,
		MimeTextHTML:  SuffixTextHTML,
		MimeTextPlain: SuffixTextPlain,
	}
)

func ToMime(s Suffix) Mime {
	m, ok := suffixToMime[s]
	if ok {
		return m
	}

	return ""
}

func ToSuffix(m Mime) Suffix {
	s, ok := mimeToSuffix[m]
	if ok {
		return s
	}

	return ""
}
