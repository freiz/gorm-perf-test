**Task**: Find out the *fastest* solution to insert multiple records to MySQL with the popular go orm framework: gorm.

# Background
## GORM
There are two versions of GORM
* [GORM V1](https://v1.gorm.io/), looks like discontinued.
* [GORM V2](https://gorm.io/), new version of this library with breaking changes.

## Insertion Method
### Sequential
One thread, insert record one by one.
### MultiThread
Multiple threads, each thread insert record one by one.
### Batch
One thread, group records into batch, then insert batch. Note: V1 does not have batch insertion functionality. 
I used one popular batch extension here, however, the implementation cannot insert association.

# Experiment
## Experiment Setting
I used local MySQL. Before any test run, all the data will be flushed. The test data models are faked. 

* The simple model only contains primitive type properties, so maps to 1 DB table
* The complex model contains a one to many mapping of another simple model, so maps to 2 DB tables

All the test data are faked also.

The environment is not exactly like production
* The service talks to local MySQL instance, the network roundtrip is minimized
* The local MySQL does not have historical data, so should be faster than the real production instance
However relative performance comparisons between different settings are still useful qualitatively.

## Experiment Result
| Data Model | GORM Version | Sequential (s) | MultiThread (s) | Batch (s) |
| --- | --- | --- | --- | --- |
| Simple | V1 | 9.7 | 11.7 | 1.5 |
| Simple | V2 | 9.5 | 10.0 | **1.0** |
| Complex | V1 | 21.0 | 20.8 | NA (Not Support) |
| Complex | V2 | 2.7 | 2.7 | **2.2** |

# Conclusions
* The MultiThread method does not outperform the Sequential method.
* Batching is the key to speed up. (for Complex.V2 however, Sequential performed as well as Batch, I think it is because in the mocked data, 1 person record have around 50 credit card record, I'm guessing the V2 implmentation batched all these SQL statements together while V1 not)
* V2 is better in this scenario. V1 does not support batching for Complex model

