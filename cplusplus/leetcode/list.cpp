

ListNode* reverseList(ListNode* head) {
        ListNode * P = NULL;
        ListNode * M =head;
        ListNode * L = head->next;
        while(M != NULL && L!= NULL) {
            M->next = P;
            P = M;
            M = L;
            L = L->next;
        }
        head = P;
        return head;
    }
