import { Queue } from "~/utils/queue";
import { sleep } from "~/utils/index";
import * as _ from "lodash-es";

export interface ProgressStatus {
  importing: boolean;
  total: number;
  done: number;
}

export interface JobStatus {
  isSuccess: boolean;
  error?: object | string;
}

interface IJob {
  start(): Promise<JobStatus>;
}

export class Job<T> implements IJob {
  private job?: Promise<JobStatus>;
  private running: boolean = false;
  private finished: boolean = false;
  public result?: JobStatus;
  public id: string;

  constructor(
    private handlerFunction: (data: T) => Promise<JobStatus>,
    private data: T,
    private debug: boolean = false,
  ) {
    this.id = `job-${_.random(100000, 1000000)}`;
  }

  async start(): Promise<JobStatus> {
    if (!this.running || !this.job) {
      if (this.debug) console.log(`Starting job ${this.id}...`);
      this.running = true;
      this.job = this.handlerFunction(this.data);
    }

    return this.job!.then((result) => {
      this.result = result;
      return this.result;
    })
      .catch((error) => {
        this.result = {
          isSuccess: false,
          error: error,
        };
        return this.result;
      })
      .finally(() => {
        this.finished = true;
        if (this.debug) console.log(`Finish job ${this.id}!`);
        return this.result;
      });
  }
}

interface IErrorJob<T> {
  data: T;
  error?: string | object;
}

interface IJobHandler<T> {
  start(): Promise<any>;

  addJob(item: T): void;

  size(): number;

  getSuccessJobs(): T[];

  getErrorJobs(): IErrorJob<T>[];

  isFinish(): boolean;

  isRunning(): boolean;
}

export interface JobHandlerOption {
  timeoutEachTask?: number | undefined;
  timeout?: number | undefined;
  waitTime?: number;
  workerLimit?: number;
  debug?: boolean;
  finishText?: string;
}

export class JobHandler<T> implements IJobHandler<T> {
  private _queue: Queue<T> = new Queue<T>();
  private _successJobs: T[] = [];
  private _errorJobs: IErrorJob<T>[] = [];
  private _total = 0;
  private _done = 0;
  private _jobs: Job<T>[] = [];
  readonly _options: JobHandlerOption = {
    timeoutEachTask: Infinity,
    timeout: Infinity,
    waitTime: 100,
    workerLimit: 10,
    debug: false,
    finishText: $i18n.t("crud.finished").toString(),
  };
  private _finished: boolean = false;
  private _running: boolean = false;
  readonly handlerFunction: (data: T) => Promise<JobStatus>;
  readonly updateProgress: () => void;
  public importProgress?: ProgressStatus;

  constructor(
    handlerFn: (data: T) => Promise<JobStatus>,
    importProgress?: ProgressStatus,
    data?: any[],
    opts?: JobHandlerOption,
  ) {
    this._options = {
      ...this._options,
      ...opts,
    };
    this.handlerFunction = handlerFn;
    this.importProgress = importProgress;
    this.updateProgress = () => {
      if (importProgress) {
        importProgress.total = this._total;
        importProgress.done = this._done;
      }
    };
    if (data && data.length > 0) {
      data.map(async (rowData) => {
        this.addJob(rowData);
      });
    }
  }

  size(): number {
    return this._queue.size();
  }

  async start(): Promise<any> {
    if (!this.isRunning()) {
      if (this._options.debug) console.log("Starting...");
      this._running = true;
      this._done = 0;
      this.updateProgress();
      this._finished = false;
      this._jobs = [];
      while (this._running && !this._finished) {
        if (this.importProgress && !this.importProgress?.importing) {
          console.log("Force stop importing!");
          break;
        }
        if (this._queue.size() === 0) {
          await sleep(this._options.waitTime!);
          if (this._queue.size() === 0) {
            if (this._options.debug) console.log("Timeout for waiting new task!");
            break;
          }
        } else {
          if (!this._options.workerLimit || this._jobs.length < this._options.workerLimit) {
            const data = this._queue.dequeue();
            if (data !== undefined) {
              const job = new Job<T>(this.handlerFunction, data);
              this._jobs.push(job);
              await job
                .start()
                .then((result) => {
                  if (result.isSuccess) this._successJobs.push(data!);
                  else
                    this._errorJobs.push({
                      error: result.error,
                      data: data,
                    } as IErrorJob<T>);
                })
                .finally(() => {
                  this._done++;
                  this.updateProgress();
                  this._jobs.splice(this._jobs.indexOf(job), 1);
                });
            }
          } else {
            await sleep(1000);
            if (this._options.debug) console.log("Reach worker limit! Must wait!");
          }
        }
      }

      return this.waitFinish().then((result: any) => {
        // Vue.notify({
        //   type: "success",
        //   title: $i18n.t("notification.success_title").toString(),
        //   text: this._options.finishText,
        //   duration: NOTIFICATION_DURATION,
        // });
        return result;
      });
    }
  }

  async waitFinish() {
    while (this._jobs.length > 0) {
      await sleep(500);
    }
    this._running = false;
    this._finished = true;
    if (this._options.debug) console.log("FINISHED");
    return this;
  }

  addJob(item: T): void {
    this._total++;
    this._queue.enqueue(item);
    this.updateProgress();
    if (this._options.debug) console.log("New job added!", this._total);
  }

  getSuccessJobs(): T[] {
    return this._successJobs;
  }

  getErrorJobs(): IErrorJob<T>[] {
    return this._errorJobs;
  }

  isFinish(): boolean {
    return this._finished;
  }

  isRunning(): boolean {
    return this._running;
  }
}

export const wrappedAxiosHandler = (promise: Promise<any>): Promise<JobStatus> => {
  return promise
    .then((resp) => {
      if (resp.data.success) return { isSuccess: true } as JobStatus;
      else
        return {
          isSuccess: false,
          error: resp.data.message || resp.data.error,
        } as JobStatus;
    })
    .catch((error) => {
      return { isSuccess: false, error: error } as JobStatus;
    });
};
