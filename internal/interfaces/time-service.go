package interfaces

type TimeServiceInterface interface {
    GetCurrentTime() (string, error)
    GetGreeting() (string, error)
    GetTimeWithFormat(format string) (string, error)
    CalculateTimeDifference(from, to string) (string, error)
}