 /*
 Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

 Licensed under the Apache License, Version 2.0 (the "License"). You may
 not use this file except in compliance with the License. A copy of the
 License is located at

	http://aws.amazon.com/apache2.0/

 or in the "license" file accompanying this file. This file is distributed
 on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 express or implied. See the License for the specific language governing
 permissions and limitations under the License.
 */

#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <sys/sem.h>

int createdSemId;

static void sigdown(int signo) {
  psignal(signo, "Shutting down, got signal");
  semctl(createdSemId, 0, IPC_RMID);
  exit(0);
}

int main() {
  if (getpid() != 1)
    /* Not an error because pause sees use outside of infra containers. */
  fprintf(stderr, "Warning: pause should be the first process\n");
  if (sigaction(SIGINT, &(struct sigaction){.sa_handler = sigdown}, NULL) < 0)
    return 1;
  if (sigaction(SIGTERM, &(struct sigaction){.sa_handler = sigdown}, NULL) < 0)
    return 2;

  key_t key = 5;
  int semflg = IPC_CREAT | 0644;
  int nsems = 1;

  if ((createdSemId = semget(key, nsems, semflg)) == -1) {
    perror("semget: semget failed. Unable to create IPC semaphore");
    exit(1);
  }
  
  //int i;
  for (;;) {
      sleep(1);
  }
  return 1;
}
