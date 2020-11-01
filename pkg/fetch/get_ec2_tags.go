/*
 * Copyright 2020 Taylor McClure
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package fetch

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

type smdInstance struct {
	InstanceID   string
	CreationTime int64
	SmdDuration  int64
	SmdTerminate bool
}

func GetSmdInstances(session *session.Session) []*ec2.Reservation {
	svc := ec2.New(session)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:smd"),
				Values: []*string{
					aws.String("true"),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Error(aerr.Error())
			}
		} else {
			log.Error(err.Error())
		}
	}

	//smdInstances := []smdInstance

	for _, r := range result.Reservations {
		// fmt.Println(i)
		// fmt.Println(i2)
		for _, ins := range r.Instances {
			/*
				fmt.Println(ins.Tags)
				fmt.Println(*ins.InstanceId)
				fmt.Println(ins.LaunchTime.Unix())
			*/
			for _, k := range ins.Tags {
				fmt.Println(k)
				if *k.Key == "smd" {
					fmt.Println(*k.Value)
					if *k.Value == "false" {
						log.Info("smd tag was set, but to false. Dropping this instance from evaluation.")
						break
					}
				}
			}
			// duration := tags["smd_duration"]
			// append(smdInstances, smdInstance{InstanceID: *ins.InstanceId, CreationTime: nil, SmdDuration: ins.*Tags[0]["smd_duration"]})

		}
	}
	return result.Reservations
}
