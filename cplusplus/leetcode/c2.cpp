
 #include<iostream>
 #include<vector>

 using namespace std;

 struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(NULL) {}
 };

void mergeKLists(vector<ListNode*>& lists) {
        
        ListNode *head = new ListNode(1);
        ListNode *resp = new ListNode(-1);
        //while(1) {
	int x = 9;
	while(1){
            int min = 1 << 30;
		std:cout << min << std::endl;
            int min_idx = -1;
            for(int i = 0 ; i < lists.size(); i++) {
                if(lists[i] != NULL) {
                    if (lists[i]->val < min) {
                        min = lists[i]->val;
                        min_idx = i;
			std::cout << lists[i]->val << " " << min << std::endl;
                    }
                }
            }
			std::cout << min_idx << std::endl;
            
            if(min_idx == -1){
                break;
            }

	    if( resp == NULL) {
	
		resp = lists[min_idx];
		std::cout << "9 好" << std::endl;
		std::cout << resp->val << std::endl;
	    }
            
            head->next = lists[min_idx];
            lists[min_idx] = lists[min_idx]->next;
	    if(lists[0] != NULL)
	    std::cout << " tiantian 0 is " << lists[0]->val << std::endl;
	    if(lists[1] != NULL)
	    std::cout << " tiantian 1 is " << lists[1]->val << std::endl;
	    if(lists[2] != NULL)
	    std::cout << " tiantian 2 is " << lists[2]->val << std::endl;
	
	    if (head != NULL)
            	head = head->next;
	    x--;
        }

        
   //     return resp;
        
    }

   void test(vector<ListNode *> nums) {
	
	nums[0] = nums[0]->next;
	std::cout<< nums[0]->val << std::endl;
	nums[0] = nums[0]->next;
	std::cout<< nums[0]->val << std::endl;
   }

   int main() {
	ListNode *l1 = new ListNode(1);
	ListNode *l2 = new ListNode(4);
	ListNode *l3 = new ListNode(5);

	l1->next = l2; l2->next = l3; l3->next = NULL;

	ListNode *n1 = new ListNode(1);
	ListNode *n2 = new ListNode(3);
	ListNode *n3 = new ListNode(4);

	n1->next = n2; n2->next = n3; n3->next = NULL;

	ListNode *m1 = new ListNode(2);
	ListNode *m2 = new ListNode(6);

	m1->next = m2; m2->next = NULL;

	vector<ListNode *> nums ;
	nums.push_back(l1);
	nums.push_back(n1);
	nums.push_back(m1);

	mergeKLists(nums);
//	test(nums);

	delete l1;
	delete l2;
	delete l3;

	delete n1;
	delete n2;
	delete n3;

	delete m1;
	delete m2;
   }
