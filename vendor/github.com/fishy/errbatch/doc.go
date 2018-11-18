// Package errbatch provides ErrBatch, which can be used to compile multiple
// errors into a single error.
//
// An example of how to use it in your functions:
//
//     type worker func() error
//
//     func runWorksParallel(works []worker) error {
//         errChan := make(chan error, len(works))
//         var wg sync.WaitGroup
//         wg.Add(len(works))
//
//         for _, work := range works {
//             go func(work worker) {
//                 errChan <- work()
//             }(work)
//         }
//
//         wg.Wait()
//         var batch errbatch.ErrBatch
//         for err := range errChan {
//             // nil errors will be auto skipped
//             batch.Add(err)
//         }
//         // If all works succeeded, Compile() returns nil.
//         // If only one work failed, Compile() returns that error directly
//         // instead of wrapping it inside ErrBatch.
//         return batch.Compile()
//     }
package errbatch
