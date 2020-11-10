package redis

import (
	"encoding/json"
	"fmt"
)

const (
	// ExpireSecondsOption time
	ExpireSecondsOption = "EX"

	// ExpireMillisecondsOption time
	ExpireMillisecondsOption = "PX"

	// ExistOption = is exist key
	ExistOption = "XX"

	// NotExitOption is not exist key
	NotExitOption = "NX"
)

// Options is redis options
type Options struct {
	key               string
	oldKey            string
	keyValue          string
	data              interface{}
	isCount           bool
	count             int // for lrem
	expireSecond      int64
	expireMilliSecond int64
	notExist          bool
	exist             bool
	keepTTL           int
	isRange           bool
	rangeLower        int
	rangeUpper        int
	minInf            string
	maxInf            string
}

// NewCommandOptions is build new options
func NewCommandOptions() *Options {
	return &Options{
		expireSecond:      0,
		expireMilliSecond: 0,
		keepTTL:           0,
	}
}

// CommandOptions is function options
type CommandOptions func(*Options)

func (options *Options) setKey(key string) error {
	if key != "" {
		options.key = key
		return nil
	}

	return fmt.Errorf("Error key is empty")
}

func (options *Options) setOldKey(key string) error {
	if key != "" {
		options.oldKey = key
		return nil
	}

	return fmt.Errorf("Error old key is empty")
}

func (options *Options) setKeyValue(keyValue string) error {
	if keyValue != "" {
		options.keyValue = keyValue
		return nil
	}

	return fmt.Errorf("Error key value is empty")
}

func (options *Options) setData(data interface{}) error {
	if data != nil {
		options.data = data
		return nil
	}

	return fmt.Errorf("Error data of key is empty")
}

func (options *Options) setIsCount() error {
	options.isCount = !options.isCount
	return nil
}

func (options *Options) setCount(count int) error {
	options.count = count
	return nil
}

func (options *Options) setExpireS(exp int64) error {
	if exp > -1 {
		options.expireSecond = exp
		return nil
	}

	return fmt.Errorf("Error expired of key is empty")
}

func (options *Options) setExpireM(exp int64) error {
	if exp > -1 {
		options.expireMilliSecond = exp
		return nil
	}

	return fmt.Errorf("Error expired of key is empty")
}

func (options *Options) setExist(exist bool) error {
	options.exist = exist

	return nil
}

func (options *Options) setNotExist(notExist bool) error {
	options.notExist = notExist

	return nil
}

func (options *Options) setIsRange() error {
	options.isRange = true

	return nil
}

func (options *Options) setRange(lower, upper int) error {
	options.rangeUpper = upper
	options.rangeLower = lower

	return nil
}

// Build is generate options
func (options *Options) Build() ([]interface{}, error) {
	var op []interface{}
	if options == nil {
		return op, nil
	}

	if options.exist && options.notExist {
		return nil, fmt.Errorf("Error invalid. Please choose exist or not exist, dont take all")
	}

	if options.expireSecond > 0 && options.expireMilliSecond > 0 {
		return nil, fmt.Errorf("Error invalid. Please choose expire with second or expire with millisecond, dont take all")
	}

	if options.oldKey != "" {
		op = append(op, options.oldKey)
	}

	if options.key != "" {
		op = append(op, options.key)
	}

	if options.isCount {
		op = append(op, options.count)
	}

	if options.keyValue != "" {
		op = append(op, options.keyValue)
	}

	if options.data != nil {
		raw, err := json.Marshal(options.data)
		if err != nil {
			return op, err
		}

		op = append(op, raw)
	}

	if options.expireSecond > 0 {
		op = append(op, ExpireSecondsOption)
		op = append(op, options.expireSecond)
	}

	if options.expireMilliSecond > 0 {
		op = append(op, ExpireMillisecondsOption)
		op = append(op, options.expireMilliSecond)
	}

	if options.exist {
		op = append(op, ExistOption)
	}

	if options.notExist {
		op = append(op, NotExitOption)
	}

	if options.isRange {
		op = append(op, options.rangeLower)
		op = append(op, options.rangeUpper)
	}

	if options.minInf != "" {
		op = append(op, options.minInf)
	}

	if options.maxInf != "" {
		op = append(op, options.maxInf)
	}

	//return fmt.Sprintf("%s", strings.Join(op, ",")), nil
	return op, nil
}

// WithKey is set key
func WithKey(key string) CommandOptions {
	return func(options *Options) {
		options.setKey(key)
	}
}

// WithOldKey is old key
func WithOldKey(key string) CommandOptions {
	return func(options *Options) {
		options.setOldKey(key)
	}
}

// WithKeyValue is key value
func WithKeyValue(keyValue string) CommandOptions {
	return func(options *Options) {
		options.setKeyValue(keyValue)
	}
}

// WithCount is count argument
func WithCount(count int) CommandOptions {
	return func(options *Options) {
		options.setIsCount()
		options.setCount(count)
	}
}

// WithData is data of key
func WithData(data interface{}) CommandOptions {
	return func(options *Options) {
		options.setData(data)
	}
}

// WithExpireSecond set expire time
func WithExpireSecond(exp int64) CommandOptions {
	return func(options *Options) {
		options.setExpireS(exp)
	}

}

// WithExpireMillisecond is expire in millisecond
func WithExpireMillisecond(exp int64) CommandOptions {
	return func(options *Options) {
		options.setExpireM(exp)
	}
}

// WithExist is exist key
func WithExist() CommandOptions {
	return func(options *Options) {
		options.setExist(true)
	}
}

// WithNotExist is not exits key
func WithNotExist() CommandOptions {
	return func(options *Options) {
		options.setNotExist(true)
	}
}

// WithRange is not exits key
func WithRange(lower, upper int) CommandOptions {
	return func(options *Options) {
		options.setIsRange()
		options.setRange(lower, upper)
	}
}

// WithMinInf is min infinity
func WithMinInf() CommandOptions {
	return func(options *Options) {
		options.minInf = "-inf"
	}
}

// WithMaxInf is max infinity
func WithMaxInf() CommandOptions {
	return func(options *Options) {
		options.maxInf = "+inf"
	}
}
