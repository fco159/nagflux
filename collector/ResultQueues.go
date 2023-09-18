package collector

import "github.com/fco159/nagflux/data"

type ResultQueues map[data.Target]chan Printable
